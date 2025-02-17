import 'package:flutter/material.dart';

class SettingDropdown extends StatelessWidget {
  final String title;
  final String subtitle;
  final String value;
  final Function(String?) onChanged;

  const SettingDropdown({
    super.key,
    required this.title,
    required this.subtitle,
    required this.value,
    required this.onChanged,
  });

  @override
  Widget build(BuildContext context) {
    return ListTile(
      title: Text(title),
      subtitle: Text(subtitle),
      trailing: DropdownButton<String>(
        value: value,
        items: <String>[
          '15',
          '30',
          '45',
          '60',
        ].map<DropdownMenuItem<String>>((String value) {
          return DropdownMenuItem<String>(
            value: value,
            child: Text('$value'),
          );
        }).toList(),
        onChanged: onChanged,
      ),
    );
  }
}
