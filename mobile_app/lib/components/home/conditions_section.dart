import 'package:flutter/material.dart';
import 'condition_card.dart';
import '../../utils/conditions_utils.dart';

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
          children: [
            Expanded(
              child: ConditionCard(
                icon: Icons.traffic,
                title: 'Traffico',
                value: getTrafficText(TrafficLevel.high),
                color: getTrafficColor(TrafficLevel.high),
              ),
            ),
            SizedBox(width: 16),
            Expanded(
              child: ConditionCard(
                icon: Icons.cloudy_snowing,
                title: 'Meteo',
                value: getWeatherText(WeatherCondition.rainy),
                color: getWeatherColor(WeatherCondition.rainy),
              ),
            ),
          ],
        ),
      ],
    );
  }
}
