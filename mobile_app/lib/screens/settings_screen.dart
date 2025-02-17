import 'package:flutter/material.dart';
import '../components/settings/setting_switch.dart';
import '../components/settings/setting_dropdown.dart';
import '../components/settings/smart_alarm_section.dart';

class SettingsScreen extends StatefulWidget {
  const SettingsScreen({super.key});

  @override
  State<SettingsScreen> createState() => _SettingsScreenState();
}

class _SettingsScreenState extends State<SettingsScreen> {
  // Impostazioni di default
  bool _trafficEnabled = true;
  bool _weatherEnabled = true;
  bool _calendarSync = false;
  String _preparationTime = '30';

  // Dati utente di esempio
  String _userName = 'Mario Rossi';
  String _userEmail = 'mario.rossi@example.com';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Impostazioni'),
      ),
      body: ListView(
        children: [
          // Sezione Profilo
          const ListTile(
            title: Text('PROFILO',
                style: TextStyle(
                  fontSize: 14,
                  fontWeight: FontWeight.bold,
                  color: Colors.blue,
                )),
          ),
          ListTile(
            leading: const CircleAvatar(
              child: Icon(Icons.person),
            ),
            title: Text(_userName),
            subtitle: Text(_userEmail),
            trailing: IconButton(
              icon: const Icon(Icons.edit),
              onPressed: () => _showEditProfileDialog(),
            ),
          ),
          ListTile(
            leading: const Icon(Icons.password),
            title: const Text('Cambia Password'),
            onTap: () => _showChangePasswordDialog(),
          ),

          // Sezione Sveglia Intelligente
          const Divider(),
          SmartAlarmSection(
            trafficEnabled: _trafficEnabled,
            weatherEnabled: _weatherEnabled,
            onTrafficChanged: (bool value) {
              setState(() {
                _trafficEnabled = value;
              });
            },
            onWeatherChanged: (bool value) {
              setState(() {
                _weatherEnabled = value;
              });
            },
          ),

          // Sezione Personalizzazione
          const ListTile(
            title: Text('PERSONALIZZAZIONE',
                style: TextStyle(
                  fontSize: 14,
                  fontWeight: FontWeight.bold,
                  color: Colors.blue,
                )),
          ),
          SettingDropdown(
            title: 'Tempo di Preparazione',
            subtitle: 'Minuti necessari per prepararsi',
            value: _preparationTime,
            onChanged: (String? newValue) {
              setState(() {
                _preparationTime = newValue!;
              });
            },
          ),
        ],
      ),
    );
  }

  void _showEditProfileDialog() {
    final nameController = TextEditingController(text: _userName);
    final emailController = TextEditingController(text: _userEmail);

    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('Modifica Profilo'),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            TextField(
              controller: nameController,
              decoration: const InputDecoration(
                labelText: 'Nome',
                icon: Icon(Icons.person),
              ),
            ),
            const SizedBox(height: 8),
          ],
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('Annulla'),
          ),
          TextButton(
            onPressed: () {
              setState(() {
                _userName = nameController.text;
                _userEmail = emailController.text;
              });
              Navigator.pop(context);
            },
            child: const Text('Salva'),
          ),
        ],
      ),
    );
  }

  void _showChangePasswordDialog() {
    final currentPasswordController = TextEditingController();
    final newPasswordController = TextEditingController();
    final confirmPasswordController = TextEditingController();

    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('Cambia Password'),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            TextField(
              controller: currentPasswordController,
              decoration: const InputDecoration(
                labelText: 'Password Attuale',
                icon: Icon(Icons.lock),
              ),
              obscureText: true,
            ),
            const SizedBox(height: 8),
            TextField(
              controller: newPasswordController,
              decoration: const InputDecoration(
                labelText: 'Nuova Password',
                icon: Icon(Icons.lock_outline),
              ),
              obscureText: true,
            ),
            const SizedBox(height: 8),
            TextField(
              controller: confirmPasswordController,
              decoration: const InputDecoration(
                labelText: 'Conferma Password',
                icon: Icon(Icons.lock_outline),
              ),
              obscureText: true,
            ),
          ],
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('Annulla'),
          ),
          TextButton(
            onPressed: () {
              // Qui andrebbe implementata la logica di validazione
              // e aggiornamento della password
              Navigator.pop(context);
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(
                  content: Text('Password aggiornata con successo'),
                ),
              );
            },
            child: const Text('Aggiorna'),
          ),
        ],
      ),
    );
  }
}
