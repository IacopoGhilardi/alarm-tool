import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';
import '../../models/alarm.dart';

class AlarmDialog extends StatefulWidget {
  final Function(Alarm) onAlarmAdded;

  const AlarmDialog({super.key, required this.onAlarmAdded});

  @override
  State<AlarmDialog> createState() => _AlarmDialogState();
}

class _AlarmDialogState extends State<AlarmDialog> {
  DateTime _selectedDateTime = DateTime.now();
  final TextEditingController _titleController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text('Aggiungi sveglia'),
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          TextField(
            controller: _titleController,
            decoration: const InputDecoration(
              labelText: 'Titolo',
            ),
          ),
          const SizedBox(height: 16),
          SizedBox(
            height: 180,
            child: CupertinoDatePicker(
              mode: CupertinoDatePickerMode.time,
              initialDateTime: _selectedDateTime,
              onDateTimeChanged: (DateTime newDateTime) {
                setState(() {
                  _selectedDateTime = newDateTime;
                });
              },
              use24hFormat: true,
            ),
          ),
        ],
      ),
      actions: [
        TextButton(
          onPressed: () {
            final newAlarm = Alarm(
              id: DateTime.now().toString(),
              title: _titleController.text.isEmpty
                  ? "Sveglia"
                  : _titleController.text,
              time: _selectedDateTime,
            );

            widget.onAlarmAdded(newAlarm);
            Navigator.of(context).pop();
          },
          child: const Text('Aggiungi'),
        ),
      ],
    );
  }

  @override
  void dispose() {
    _titleController.dispose();
    super.dispose();
  }
}
