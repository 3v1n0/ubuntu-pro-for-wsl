// Mocks generated by Mockito 5.4.4 from annotations
// in ubuntupro/test/core/agent_monitor_test.dart.
// Do not manually edit this file.

// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'dart:async' as _i5;

import 'package:agentapi/agentapi.dart' as _i2;
import 'package:grpc/grpc.dart' as _i4;
import 'package:mockito/mockito.dart' as _i1;
import 'package:ubuntupro/core/agent_api_client.dart' as _i3;

// ignore_for_file: type=lint
// ignore_for_file: avoid_redundant_argument_values
// ignore_for_file: avoid_setters_without_getters
// ignore_for_file: comment_references
// ignore_for_file: deprecated_member_use
// ignore_for_file: deprecated_member_use_from_same_package
// ignore_for_file: implementation_imports
// ignore_for_file: invalid_use_of_visible_for_testing_member
// ignore_for_file: prefer_const_constructors
// ignore_for_file: unnecessary_parenthesis
// ignore_for_file: camel_case_types
// ignore_for_file: subtype_of_sealed_class

class _FakeUIClient_0 extends _i1.SmartFake implements _i2.UIClient {
  _FakeUIClient_0(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeSubscriptionInfo_1 extends _i1.SmartFake
    implements _i2.SubscriptionInfo {
  _FakeSubscriptionInfo_1(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeConfigSources_2 extends _i1.SmartFake implements _i2.ConfigSources {
  _FakeConfigSources_2(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

/// A class which mocks [AgentApiClient].
///
/// See the documentation for Mockito's code generation for more information.
class MockAgentApiClient extends _i1.Mock implements _i3.AgentApiClient {
  MockAgentApiClient() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i2.UIClient Function(_i4.ClientChannel) get stubFactory =>
      (super.noSuchMethod(
        Invocation.getter(#stubFactory),
        returnValue: (_i4.ClientChannel __p0) => _FakeUIClient_0(
          this,
          Invocation.getter(#stubFactory),
        ),
      ) as _i2.UIClient Function(_i4.ClientChannel));

  @override
  _i5.Stream<_i3.ConnectionEvent> get onConnectionChanged =>
      (super.noSuchMethod(
        Invocation.getter(#onConnectionChanged),
        returnValue: _i5.Stream<_i3.ConnectionEvent>.empty(),
      ) as _i5.Stream<_i3.ConnectionEvent>);

  @override
  _i5.Future<bool> connectTo({
    required String? host,
    required int? port,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #connectTo,
          [],
          {
            #host: host,
            #port: port,
          },
        ),
        returnValue: _i5.Future<bool>.value(false),
      ) as _i5.Future<bool>);

  @override
  _i5.Future<_i2.SubscriptionInfo> applyProToken(String? token) =>
      (super.noSuchMethod(
        Invocation.method(
          #applyProToken,
          [token],
        ),
        returnValue:
            _i5.Future<_i2.SubscriptionInfo>.value(_FakeSubscriptionInfo_1(
          this,
          Invocation.method(
            #applyProToken,
            [token],
          ),
        )),
      ) as _i5.Future<_i2.SubscriptionInfo>);

  @override
  _i5.Future<void> applyLandscapeConfig(String? config) => (super.noSuchMethod(
        Invocation.method(
          #applyLandscapeConfig,
          [config],
        ),
        returnValue: _i5.Future<void>.value(),
        returnValueForMissingStub: _i5.Future<void>.value(),
      ) as _i5.Future<void>);

  @override
  _i5.Future<bool> ping() => (super.noSuchMethod(
        Invocation.method(
          #ping,
          [],
        ),
        returnValue: _i5.Future<bool>.value(false),
      ) as _i5.Future<bool>);

  @override
  _i5.Future<_i2.ConfigSources> configSources() => (super.noSuchMethod(
        Invocation.method(
          #configSources,
          [],
        ),
        returnValue: _i5.Future<_i2.ConfigSources>.value(_FakeConfigSources_2(
          this,
          Invocation.method(
            #configSources,
            [],
          ),
        )),
      ) as _i5.Future<_i2.ConfigSources>);

  @override
  _i5.Future<_i2.SubscriptionInfo> notifyPurchase() => (super.noSuchMethod(
        Invocation.method(
          #notifyPurchase,
          [],
        ),
        returnValue:
            _i5.Future<_i2.SubscriptionInfo>.value(_FakeSubscriptionInfo_1(
          this,
          Invocation.method(
            #notifyPurchase,
            [],
          ),
        )),
      ) as _i5.Future<_i2.SubscriptionInfo>);
}
