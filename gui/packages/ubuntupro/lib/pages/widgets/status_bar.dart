import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:provider/provider.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:yaru/yaru.dart';

import '/constants.dart' as constants;
import '/core/agent_connection.dart';

class StatusBar extends StatelessWidget {
  const StatusBar({
    super.key,
    this.showAgentStatus = true,
    this.launchUrlFn = launchUrl,
  });

  final bool showAgentStatus;
  final Future<bool> Function(Uri) launchUrlFn;

  static const bugIcon = Icons.bug_report_outlined;
  static const agentConnIcon = Icons.circle_rounded;

  @override
  Widget build(BuildContext context) {
    final lang = AppLocalizations.of(context);
    return Row(
      children: <Widget>[
        const SizedBox(width: 8.0),
        Text(
          constants.kVersion,
          style: Theme.of(context).textTheme.bodySmall?.copyWith(
                color: YaruColors.warmGrey,
              ),
        ),
        const Spacer(),
        IconButton(
          tooltip: lang.statusBarReportBugTooltip,
          onPressed: () {
            launchUrlFn(
              Uri.https(
                'github.com',
                '/canonical/ubuntu-pro-for-wsl/issues/new',
                {'labels': 'bug', 'template': 'bug_report.yml'},
              ),
            );
          }, // open link to new issue in GH
          icon: const Icon(
            bugIcon,
            color: YaruColors.warmGrey,
            size: 14.0,
          ),
        ),
        if (showAgentStatus)
          Consumer<AgentConnection>(
            builder: (context, conn, _) => IconButton(
              tooltip: conn.isConnected
                  ? lang.statusBarAgentRunningTooltip
                  : lang.statusBarAgentDownTooltip,
              icon: Icon(
                size: 14.0,
                agentConnIcon,
                color: conn.isConnected
                    ? YaruColors.of(context).success
                    : YaruColors.of(context).error,
              ),
              onPressed: conn.isConnected ? null : conn.restartAgent,
            ),
          ),
      ],
    );
  }
}
