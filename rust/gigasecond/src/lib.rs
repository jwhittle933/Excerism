use chrono::{DateTime, Duration, Utc};

const GIGA: i64 = 1_000_000_000;

// Returns a Utc DateTime one billion seconds after start.
// let off: DateTime<Utc> = ;
pub fn after(start: DateTime<Utc>) -> DateTime<Utc> {
    start.checked_add_signed(Duration::seconds(GIGA)).unwrap()
}
