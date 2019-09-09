use chrono::{DateTime, Duration, Utc};

// Returns a Utc DateTime one billion seconds after start.
// let off: DateTime<Utc> = ;
pub fn after(start: DateTime<Utc>) -> DateTime<Utc> {
    start.checked_add_signed(Duration::seconds(1000000000)).unwrap()
}
