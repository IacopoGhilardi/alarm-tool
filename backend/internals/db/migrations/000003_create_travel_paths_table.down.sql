CREATE TABLE IF NOT EXISTS travel_paths (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    user_id INTEGER NOT NULL REFERENCES users(id),
    start_location JSONB NOT NULL,
    end_location JSONB NOT NULL
);

CREATE INDEX idx_travel_paths_user_id ON travel_paths(user_id);