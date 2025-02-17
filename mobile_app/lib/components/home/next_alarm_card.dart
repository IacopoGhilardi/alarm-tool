import 'package:flutter/material.dart';
import 'package:url_launcher/url_launcher.dart';

class NextAlarmCard extends StatefulWidget {
  const NextAlarmCard({super.key});

  @override
  State<NextAlarmCard> createState() => _NextAlarmCardState();
}

class _NextAlarmCardState extends State<NextAlarmCard> {
  bool _isAlarmActive = true;

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const Text(
              'Prossima Sveglia',
              style: TextStyle(
                fontSize: 18,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 8),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                const Text(
                  '07:30',
                  style: TextStyle(
                    fontSize: 36,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                Switch(
                  value: _isAlarmActive,
                  onChanged: (value) {
                    setState(() {
                      _isAlarmActive = value;
                    });
                  },
                ),
              ],
            ),
            const Text('Tra 8 ore e 23 minuti'),
            const SizedBox(height: 16),
            ElevatedButton.icon(
              onPressed: () async {
                final Uri uri = Uri.parse(
                    'https://www.google.com/maps/dir/?api=1&destination=DESTINAZIONE');
                if (await canLaunchUrl(uri)) {
                  await launchUrl(uri);
                } else {
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(
                      content: Text('Impossibile aprire le indicazioni'),
                    ),
                  );
                }
              },
              icon: const Icon(Icons.directions),
              label: const Text('Ottieni Indicazioni'),
              style: ElevatedButton.styleFrom(
                minimumSize: const Size(double.infinity, 36),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
