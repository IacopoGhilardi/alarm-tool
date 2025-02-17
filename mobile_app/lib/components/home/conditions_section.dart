import 'package:flutter/material.dart';
import 'condition_card.dart';

class ConditionsSection extends StatelessWidget {
  const ConditionsSection({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const Text(
          'Condizioni Attuali',
          style: TextStyle(
            fontSize: 18,
            fontWeight: FontWeight.bold,
          ),
        ),
        const SizedBox(height: 16),
        Row(
          children: const [
            Expanded(
              child: ConditionCard(
                icon: Icons.traffic,
                title: 'Traffico',
                value: 'Moderato',
                color: Colors.orange,
              ),
            ),
            SizedBox(width: 16),
            Expanded(
              child: ConditionCard(
                icon: Icons.cloud,
                title: 'Meteo',
                value: 'Sereno',
                color: Colors.blue,
              ),
            ),
          ],
        ),
      ],
    );
  }
}
