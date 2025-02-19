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
  final TextEditingController _destinationController = TextEditingController();
  bool _snoozeEnabled = true;

  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      navigationBar: CupertinoNavigationBar(
        border: null,
        leading: CupertinoButton(
          padding: EdgeInsets.zero,
          child: const Text('Annulla'),
          onPressed: () => Navigator.of(context).pop(),
        ),
        middle: const Text(
          'Aggiungi sveglia',
          style: TextStyle(color: CupertinoColors.white),
        ),
        trailing: CupertinoButton(
          padding: EdgeInsets.zero,
          child: const Text('Salva'),
          onPressed: () {
            final newAlarm = Alarm(
              id: DateTime.now().toString(),
              title: _titleController.text.isEmpty
                  ? "Sveglia"
                  : _titleController.text,
              time: _selectedDateTime,
              destination: _destinationController.text,
            );
            widget.onAlarmAdded(newAlarm);
            Navigator.of(context).pop();
          },
        ),
      ),
      child: Container(
        child: SafeArea(
          child: Column(
            children: [
              SizedBox(
                height: 150,
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
              CupertinoListSection.insetGrouped(
                children: [
                  CupertinoListTile(
                    title: const Text('Titolo'),
                    trailing: SizedBox(
                      width: 150,
                      child: CupertinoTextField.borderless(
                        controller: _titleController,
                        placeholder: 'Alarm',
                        style: const TextStyle(color: CupertinoColors.black),
                        textAlign: TextAlign.right,
                      ),
                    ),
                  ),
                  CupertinoListTile(
                    title: const Text('Destinazione'),
                    trailing: SizedBox(
                      width: 150,
                      child: CupertinoTextField.borderless(
                        controller: _destinationController,
                        placeholder: 'Address',
                        style: const TextStyle(color: CupertinoColors.white),
                        textAlign: TextAlign.right,
                      ),
                    ),
                  ),
                  const CupertinoListTile(
                    title: Text('Ripetizione'),
                    trailing: Text(
                      'Mai',
                      style: TextStyle(color: CupertinoColors.systemGrey),
                    ),
                  ),
                  CupertinoListTile(
                    title: const Text('Sveglia'),
                    trailing: CupertinoSwitch(
                      value: _snoozeEnabled,
                      onChanged: (bool value) {
                        setState(() {
                          _snoozeEnabled = value;
                        });
                      },
                    ),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }

  @override
  void dispose() {
    _titleController.dispose();
    _destinationController.dispose();
    super.dispose();
  }
}
