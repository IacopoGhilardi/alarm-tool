import 'package:flutter/material.dart';
import 'setting_switch.dart';

class SmartAlarmSection extends StatelessWidget {
  final bool trafficEnabled;
  final bool weatherEnabled;
  final Function(bool) onTrafficChanged;
  final Function(bool) onWeatherChanged;

  const SmartAlarmSection({
    super.key,
    required this.trafficEnabled,
    required this.weatherEnabled,
    required this.onTrafficChanged,
    required this.onWeatherChanged,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        const ListTile(
          title: Text(
            'SVEGLIA INTELLIGENTE',
            style: TextStyle(
              fontSize: 14,
              fontWeight: FontWeight.bold,
              color: Colors.blue,
            ),
          ),
        ),
        SettingSwitch(
          title: 'Monitoraggio Traffico',
          subtitle: 'Regola la sveglia in base al traffico',
          value: trafficEnabled,
          onChanged: onTrafficChanged,
        ),
        SettingSwitch(
          title: 'Monitoraggio Meteo',
          subtitle: 'Regola la sveglia in base al meteo',
          value: weatherEnabled,
          onChanged: onWeatherChanged,
        ),
        const Divider(),
      ],
    );
  }
}
