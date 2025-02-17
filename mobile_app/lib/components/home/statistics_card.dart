import 'package:flutter/material.dart';

class StatisticsCard extends StatelessWidget {
  const StatisticsCard({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const Text(
          'Statistiche Settimanali',
          style: TextStyle(
            fontSize: 18,
            fontWeight: FontWeight.bold,
          ),
        ),
        const SizedBox(height: 16),
        Card(
          child: Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              children: const [
                StatisticRow(label: 'Sveglie in orario', value: '5/7'),
                Divider(),
                StatisticRow(
                    label: 'Tempo medio di preparazione', value: '35 min'),
                Divider(),
                StatisticRow(label: 'Risparmio tempo', value: '15 min'),
              ],
            ),
          ),
        ),
      ],
    );
  }
}

class StatisticRow extends StatelessWidget {
  final String label;
  final String value;

  const StatisticRow({
    super.key,
    required this.label,
    required this.value,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 8.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Text(label),
          Text(
            value,
            style: const TextStyle(
              fontWeight: FontWeight.bold,
            ),
          ),
        ],
      ),
    );
  }
}
