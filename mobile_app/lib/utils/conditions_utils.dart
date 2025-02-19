import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

enum TrafficLevel {
  low,
  moderate,
  high,
  severe,
}

enum WeatherCondition {
  sunny,
  cloudy,
  rainy,
}

Color getTrafficColor(TrafficLevel level) {
  switch (level) {
    case TrafficLevel.low:
      return Colors.green;
    case TrafficLevel.moderate:
      return Colors.orange;
    case TrafficLevel.high:
      return Colors.deepOrange;
    case TrafficLevel.severe:
      return Colors.red;
  }
}

String getTrafficText(TrafficLevel level) {
  switch (level) {
    case TrafficLevel.low:
      return 'Basso';
    case TrafficLevel.moderate:
      return 'Moderato';
    case TrafficLevel.high:
      return 'Alto';
    case TrafficLevel.severe:
      return 'Molto Alto';
  }
}

Color getWeatherColor(WeatherCondition condition) {
  switch (condition) {
    case WeatherCondition.sunny:
      return Colors.blue;
    case WeatherCondition.cloudy:
      return Colors.grey;
    case WeatherCondition.rainy:
      return Colors.blueGrey;
  }
}

String getWeatherText(WeatherCondition condition) {
  switch (condition) {
    case WeatherCondition.sunny:
      return 'Soleggiato';
    case WeatherCondition.cloudy:
      return 'Nuvoloso';
    case WeatherCondition.rainy:
      return 'Pioggia';
  }
}

IconData getWeatherIcon(WeatherCondition condition) {
  switch (condition) {
    case WeatherCondition.sunny:
      return Icons.sunny;
    case WeatherCondition.cloudy:
      return Icons.cloud;
    case WeatherCondition.rainy:
      return Icons.cloudy_snowing;
  }
}
