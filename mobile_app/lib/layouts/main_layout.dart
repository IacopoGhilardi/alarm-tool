import 'package:flutter/material.dart';
import '../models/navigation_items.dart';

class MainLayout extends StatefulWidget {
  const MainLayout({super.key});

  @override
  State<MainLayout> createState() => _MainLayoutState();
}

class _MainLayoutState extends State<MainLayout> {
  // Tiene traccia dell'indice della schermata selezionata
  int _selectedIndex = 0;

  // Metodo chiamato quando un item della bottom bar viene tappato
  void _onItemTapped(int index) {
    print('Item tapped: $index');
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      // Il body mostra la schermata corrispondente all'indice selezionato
      body: NavigationItems.screens[_selectedIndex],

      // Configurazione della bottom navigation bar
      bottomNavigationBar: BottomNavigationBar(
        // Fixed impedisce l'animazione di shift quando ci sono più di 3 items
        type: BottomNavigationBarType.fixed,

        // L'indice correntemente selezionato
        currentIndex: _selectedIndex,

        // Callback quando un item viene tappato
        onTap: _onItemTapped,

        // Lista degli items nella bottom bar
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.alarm),
            label: 'Alarms',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.settings),
            label: 'Settings',
          ),
        ],
      ),
    );
  }
}
