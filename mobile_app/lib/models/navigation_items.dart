import 'package:flutter/material.dart';
import '../screens/home_screen.dart';
import '../screens/alarm_screen.dart';
import '../screens/settings_screen.dart';
import '../models/alarm.dart';

class NavigationItems {
  static final List<Widget> screens = [
    const HomeScreen(),
    AlarmScreen(alarms: [
      Alarm(id: '1', title: 'Alarm 1', time: DateTime.now()),
      Alarm(id: '2', title: 'Alarm 2', time: DateTime.now()),
    ]),
    const SettingsScreen(),
  ];
}
