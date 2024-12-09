import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:flutter/foundation.dart' show ChangeNotifier, kDebugMode;
import 'package:grpc/grpc.dart' show GrpcError;
import 'package:pkcs7/pkcs7.dart';

import '/core/agent_api_client.dart';

/// The view model for the Landscape configuration page.
/// This class is responsible for managing the state of the Landscape configuration form, including its subforms
/// and submit the active form data when complete, disregarding the inactive ones.
/// Data validation is delegated to the subform data models.
class LandscapeModel extends ChangeNotifier {
  /// The client connection to the background agent.
  final AgentApiClient client;

  LandscapeModel(this.client);

  /// The URL to be shown in the UI.
  static final landscapeURI = Uri.https('ubuntu.com', '/landscape');

  /// Whether the current form is complete (ready to be submitted).
  bool get isComplete => _active.isComplete;

  /// The current configuration type, allowing the UI to show the correct form.
  LandscapeConfigType get configType => _current;
  LandscapeConfigType _current = LandscapeConfigType.saas;

  // The active configuration form data, a shortcut to reduce some switch statements
  // and avoid relying on ducktyping when serializing the config or checking for completeness.
  late LandscapeConfig _active = saas;

  /// The configuration form data for the SaaS configuration.
  final LandscapeSaasConfig saas = LandscapeSaasConfig();

  // TODO: Remove this condition when Landscape SaaS starts supporting WSL.
  bool get isSaaSSupported => kDebugMode;

  /// The configuration form data for the custom configuration.
  final LandscapeCustomConfig custom = LandscapeCustomConfig();

  /// Allows the UI to inform the selected configuration type.
  void setConfigType(LandscapeConfigType? value) {
    if (value == null) return;
    _current = value;
    switch (configType) {
      case LandscapeConfigType.saas:
        _active = saas;
      case LandscapeConfigType.custom:
        _active = custom;
    }

    notifyListeners();
  }

  /// Sets the registration key for the SaaS configurations.
  void setSaasRegistrationKey(String? registrationKey) {
    assert(_active is LandscapeSaasConfig);
    if (registrationKey == null) return;
    saas.registrationKey = registrationKey;
    notifyListeners();
  }

  /// Sets (and validates) the FQDN for the self-hosted configuration.
  void setFqdn(String? fqdn) {
    assert(_active is LandscapeSaasConfig);
    if (fqdn == null) return;
    saas.fqdn = fqdn;
    notifyListeners();
  }

  /// Sets (and validates) the SSL key path for the self-hosted configuration.
  void setSslKeyPath(String? sslKeyPath) {
    assert(_active is LandscapeSaasConfig);
    if (sslKeyPath == null) return;
    saas.sslKeyPath = sslKeyPath;
    notifyListeners();
  }

  /// Sets (and validates) the custom configuration path.
  void setCustomConfigPath(String? configPath) {
    assert(_active is LandscapeCustomConfig);
    if (configPath == null) return;
    custom.configPath = configPath;
    notifyListeners();
  }

  /// Translates and submits the active configuration data to the background agent, returning an error message if any.
  Future<String?> applyConfig() async {
    assert(_active.isComplete);
    final config = _active.config();
    assert(config != null);
    try {
      await client.applyLandscapeConfig(config!);
      return null;
    } on GrpcError catch (e) {
      return e.message;
    }
  }
}

/// The different types of Landscape configurations, modelled as an enum to make it easy on the UI side to switch between them.
enum LandscapeConfigType { saas, custom }

/// The alternative errors we could encounter when validating file paths submitted as part of any subform data.
enum FileError {
  notFound,
  tooLarge,
  emptyPath,
  dir,
  emptyFile,
  none,
  invalidFormat,
}

const landscapeSaas = 'landscape.canonical.com';
const validCertExtensions = ['cer', 'crt', 'der', 'pem'];

/// The base class for the closed set of Landscape configuration form types.
sealed class LandscapeConfig {
  /// Whether the form has enough data for submission.
  bool get isComplete;

  /// The raw representation of the configuration (as expected by the background agent).
  String? config();
}

/// The SaaS configuration form data: only the FQDN is mandatory.
class LandscapeSaasConfig extends LandscapeConfig {
  String _fqdn = '';
  String get fqdn => _fqdn;
  bool _fqdnError = false;
  bool get fqdnError => _fqdnError;

  String registrationKey = '';

  String _sslKeyPath = '';
  String get sslKeyPath => _sslKeyPath;

  FileError _fileError = FileError.none;
  FileError get fileError => _fileError;

  // FQDN must be a valid URL (without an explicit port) and must not be the Landscape SaaS URL.
  bool _validateFQDN(String value) {
    final uri = Uri.tryParse(value);
    _fqdnError = value.isEmpty || uri == null || uri.hasPort;

    return !_fqdnError;
  }

  /// Ensure the FQDN is a valid URL, enforcing https without requiring the user to type it.
  set fqdn(String value) {
    if (value.isNotEmpty && !value.startsWith('https://')) {
      value = 'https://$value';
    }
    if (value == _fqdn) {
      return;
    }
    if (_validateFQDN(value)) {
      _fqdn = value;
    }
  }

  // If a path is provided, then it must exist and be a non-empty file.
  bool _validatePath(String path) {
    // Empty paths are allowed, since this field is optional.
    if (path.isEmpty) {
      _fileError = FileError.none;
      return true;
    }

    final file = File(path);
    final fileStat = file.statSync();

    if (fileStat.type == FileSystemEntityType.notFound) {
      _fileError = FileError.notFound;
    } else if (fileStat.type == FileSystemEntityType.directory) {
      _fileError = FileError.dir;
    } else if (fileStat.size == 0) {
      _fileError = FileError.emptyFile;
    } else if (validCertExtensions.every((e) => !file.path.endsWith(e))) {
      _fileError = FileError.invalidFormat;
    } else if (!_validateCertificate(file)) {
      _fileError = FileError.invalidFormat;
    } else {
      _fileError = FileError.none;
    }

    return _fileError == FileError.none;
  }

  bool _validateCertificate(File file) {
    final content = file.readAsBytesSync();

    try {
      X509.fromDer(content);
      return true;
      // Various Exception or Errors can occur here when attempting a parse
      // ignore: avoid_catches_without_on_clauses
    } catch (_) {
      try {
        X509.fromPem(utf8.decode(content));
        return true;
      } on Exception catch (_) {
        return false;
      }
    }
  }

  set sslKeyPath(String value) {
    if (_validatePath(value)) {
      _sslKeyPath = value;
    }
  }

  @override
  bool get isComplete =>
      !fqdnError && fqdn.isNotEmpty && fileError == FileError.none;

  @override
  String? config() {
    if (!isComplete) return null;
    final uri = Uri.parse(_fqdn);
    final sslKeyLine = sslKeyPath.isEmpty ? '' : 'ssl_public_key = $sslKeyPath';
    final registrationKeyLine =
        registrationKey.isEmpty ? '' : 'registration_key = $registrationKey';

    return '''
[host]
url = ${uri.replace(port: 6554).authority}
[client]
url = ${uri.replace(path: '/message-system')}
ping_url = ${uri.replace(scheme: 'http').replace(path: '/ping')}
log_level = info
$sslKeyLine
$registrationKeyLine
'''
        .trimRight();
  }
}

/// The custom configuration form data: the only field available is the path to the configuration file.
class LandscapeCustomConfig extends LandscapeConfig {
  String _configPath = '';
  String get configPath => _configPath;
  FileError _fileError = FileError.none;
  FileError get fileError => _fileError;

  // The provided path must exist and be a non-empty file with bounded size.
  bool _validatePath(String path) {
    if (path.isEmpty) {
      _fileError = FileError.emptyPath;
      return false;
    }

    final fileStat = File(path).statSync();
    if (fileStat.type == FileSystemEntityType.notFound) {
      _fileError = FileError.notFound;
    } else if (fileStat.type == FileSystemEntityType.directory) {
      _fileError = FileError.dir;
    } else if (fileStat.size == 0) {
      _fileError = FileError.emptyFile;
    } else if (fileStat.size >= 1024 * 1024) {
      _fileError = FileError.tooLarge;
    } else {
      _fileError = FileError.none;
    }

    return _fileError == FileError.none;
  }

  set configPath(String value) {
    if (_configPath == value) {
      return;
    }
    if (_validatePath(value)) {
      _configPath = value;
    }
  }

  @override
  bool get isComplete => fileError == FileError.none && configPath.isNotEmpty;

  @override
  String? config() {
    if (!isComplete) return null;
    final file = File(configPath);
    return file.readAsStringSync();
  }
}
