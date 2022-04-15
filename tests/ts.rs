
fn check_eq(a: f64, b: f64, err: f64) -> bool {
    if a >= b + err || a <= b - err {
        false
    }
    else {
        true
    }
}

#[cfg(test)]
mod tests {
    use sigma::ts::random;
    use sigma::ts::stats;

    use crate::check_eq;

    #[test]
    fn gwn_stats() {
        let v = random::gaussian_white_noise(10000, 2.0);
        let mean = stats::mean(&v);
        let variance = stats::variance(&v);

        assert_eq!(check_eq(mean, 0.0, 0.1), true);
        assert_eq!(check_eq(variance, 4.0, 0.2), true);
    }
}
