mod util;
use schemars::JsonSchema;
use std::time::{Duration, SystemTime};
use util::*;

#[derive(Debug, JsonSchema)]
struct MyStruct {
    duration: Duration,
    time: SystemTime,
}

#[test]
fn duration_and_systemtime() -> TestResult {
    test_default_generated_schema::<MyStruct>("duration_and_systemtime")
}
