#[cfg(test)]
mod tests {
    use sigma::ts::arima;
    use sigma::ts::stats;

    //use crate::check_eq;

    fn check_eq(a: f64, b: f64, err: f64) -> bool {
        if a >= b + err || a <= b - err {
            false
        }
        else {
            true
        }
    }

    #[test]
    fn sim_arima_only_ar() {
        let v = arima::sim_gwn(10000, &vec![0.3, 0.2], 0, &vec![], 1.0);
        assert!(check_eq(stats::mean(&v), 0.0, 0.1));
    }
}
