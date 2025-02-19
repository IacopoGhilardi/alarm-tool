import 'package:flutter/material.dart';
import '../models/alarm.dart';
import '../components/alarm/alarm_item.dart';
import '../components/alarm/alarm_dialog.dart';

class AlarmScreen extends StatefulWidget {
  const AlarmScreen({super.key, required this.alarms});

  final List<Alarm> alarms;

  @override
  State<AlarmScreen> createState() => _AlarmScreenState();
}

class _AlarmScreenState extends State<AlarmScreen> {
  late List<Alarm> _alarms;

  @override
  void initState() {
    super.initState();
    _alarms = List.from(widget.alarms);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Alarms')),
      body: ListView.builder(
        itemCount: _alarms.length,
        itemBuilder: (context, index) {
          return AlarmItem(
              alarm: _alarms[index],
              onAlarmChanged: (alarm) {
                setState(() {
                  _alarms[index] = alarm;
                });
              });
        },
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          showModalBottomSheet(
            context: context,
            isScrollControlled: true,
            builder: (context) => Container(
              height: MediaQuery.of(context).size.height * 0.7,
              child: AlarmDialog(
                onAlarmAdded: (Alarm newAlarm) {
                  setState(() {
                    _alarms.add(newAlarm);
                  });
                },
              ),
            ),
          );
        },
        child: const Icon(Icons.add),
      ),
    );
  }
}
