use std::time::Instant;

pub fn duration<F>(f: F)
where
    F: Fn() -> (),
{
    let start = Instant::now();
    f();
    let elapsed = start.elapsed();
    if elapsed.as_secs() > 0 {
        println!(
            "  Elapsed: {}.{:03}s",
            elapsed.as_secs(),
            elapsed.subsec_millis()
        );
    } else if elapsed.subsec_millis() > 0 {
        println!(
            "  Elapsed: {:.3}ms",
            elapsed.subsec_nanos() as f64 / 1_000_000.0
        );
    } else if elapsed.subsec_micros() > 0 {
        println!(
            "  Elapsed: {:.3}µs",
            elapsed.subsec_nanos() as f64 / 1_000.0
        );
    } else {
        println!(
            "  Elapsed: {}ns",
            elapsed.subsec_nanos()
        );
    }
}
