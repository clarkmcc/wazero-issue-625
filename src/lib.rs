#![allow(dead_code)]

use std::{thread::sleep, time::Duration};

#[no_mangle]
pub extern "C" fn entrypoint() {
    sleep(Duration::from_secs(1))
}