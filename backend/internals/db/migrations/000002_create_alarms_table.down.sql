CREATE TYPE repeat_type AS ENUM ('daily', 'weekly', 'custom');

CREATE TABLE IF NOT EXISTS alarms (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    user_id INTEGER NOT NULL REFERENCES users(id),
    
    label VARCHAR(255) NOT NULL,
    time TIME NOT NULL,
    is_active BOOLEAN DEFAULT true,
    
    is_repeat BOOLEAN DEFAULT false,
    repeat_days TEXT[],
    repeat_type repeat_type,
    repeat_end_time TIMESTAMP,
    repeat_start_time TIMESTAMP,
    
    sound VARCHAR(255),
    vibration BOOLEAN DEFAULT true,
    volume INTEGER CHECK (volume >= 0 AND volume <= 100),
    snooze_enabled BOOLEAN DEFAULT true,
    snooze_interval INTEGER DEFAULT 5,
    
    smart_wakeup BOOLEAN DEFAULT false,
    wakeup_window INTEGER DEFAULT 30,
    weather_aware BOOLEAN DEFAULT false,
    traffic_aware BOOLEAN DEFAULT false
);

CREATE INDEX idx_alarms_user_id ON alarms(user_id);
CREATE INDEX idx_alarms_is_active ON alarms(is_active);