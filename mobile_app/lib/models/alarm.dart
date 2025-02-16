class Alarm {
  final String id;
  final String title;
  final DateTime time;
  final bool isActive;

  const Alarm({
    required this.id,
    required this.title,
    required this.time,
    this.isActive = true,
  });

  factory Alarm.fromJson(Map<String, dynamic> json) {
    return Alarm(
      id: json['id'],
      title: json['title'],
      time: json['time'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'title': title,
      'time': time,
    };
  }

  Alarm updateWith({
    String? title,
    DateTime? time,
    bool? isActive,
  }) {
    return Alarm(
      id: id ?? this.id,
      title: title ?? this.title,
      time: time ?? this.time,
      isActive: isActive ?? this.isActive,
    );
  }
}
