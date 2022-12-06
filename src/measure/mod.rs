use std::time::Instant;

pub fn duration<F>(f: F)
where F: Fn() -> () {
    let start = Instant::now();
    f();
    let elapsed = start.elapsed();
    if elapsed.as_secs() > 0 {
        println!("  Elapsed: {}.{:09}s", elapsed.as_secs(), elapsed.subsec_nanos());
    } else {
        println!("  Elapsed: {:.6}ms", elapsed.subsec_nanos() as f64 / 1_000_000.0);
    }
}
