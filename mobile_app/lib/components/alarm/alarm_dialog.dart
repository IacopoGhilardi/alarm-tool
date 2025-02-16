import 'package:flutter/material.dart';
import '../../models/alarm.dart';

class AlarmDialog extends StatelessWidget {
  final Function(Alarm) onAlarmAdded;

  const AlarmDialog({super.key, required this.onAlarmAdded});

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text('Aggiungi sveglia'),
      content: const Text('Inserisci il titolo e la data'),
      actions: [
        TextButton(
          onPressed: () {
            // Crea un nuovo allarme
            final newAlarm = Alarm(
              id: DateTime.now().toString(), // Genera un ID univoco
              title: "Sveglia",
              time: DateTime.now(),
            );

            // Chiama la callback per aggiungere l'allarme
            onAlarmAdded(newAlarm);

            // Chiudi il dialog
            Navigator.of(context).pop();
          },
          child: const Text('Aggiungi'),
        ),
      ],
    );
  }
}
