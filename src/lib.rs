#![allow(dead_code)]

use std::{thread::sleep, time::Duration};

extern "C" fn entrypoint() {
    sleep(Duration::from_secs(1))
}