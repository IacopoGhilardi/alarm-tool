import 'package:flutter/material.dart';
import '../components/home/next_alarm_card.dart';
import '../components/home/conditions_section.dart';
import '../components/home/statistics_card.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Smart Alarm'),
        elevation: 0,
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: const [
              NextAlarmCard(),
              SizedBox(height: 24),
              ConditionsSection(),
              SizedBox(height: 24),
              StatisticsCard(),
            ],
          ),
        ),
      ),
    );
  }
}
