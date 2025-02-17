import 'package:flutter/material.dart';
import '../../models/alarm.dart';

class AlarmItem extends StatelessWidget {
  final Alarm alarm;
  final Function(Alarm) onAlarmChanged;

  AlarmItem({required this.alarm, required this.onAlarmChanged});

  @override
  Widget build(BuildContext context) {
    return ListTile(
      title: Text(alarm.title),
      subtitle: Text(
          '${alarm.time.hour}:${alarm.time.minute.toString().padLeft(2, '0')}'),
      trailing: Switch(
        value: alarm.isActive,
        onChanged: (value) {
          changeIsActive(value);
        },
      ),
    );
  }

  void changeIsActive(bool value) {
    onAlarmChanged(alarm.updateWith(isActive: value));
  }
}
