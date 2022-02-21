use statrs::distribution::Normal;
use rand::Rng;

use crate::stats;

pub fn gaussian_random_walk(n: usize, mean: f64, std: f64) -> Vec<f64> {
    let result = gaussian_white_noise(n, mean, std);
    let mut x = 0.0;
    result.iter().map(|w| { x += w; x }).collect()
}

pub fn gaussian_white_noise(n: usize, mean: f64, std: f64) -> Vec<f64> {
    let normal = Normal::new(mean, std).unwrap();
    rand::thread_rng().sample_iter(&normal).take(n).collect()
}

pub fn diff(v: &[f64], lag: usize) -> Vec<f64> {
    let n= v.len();
    let mut result: Vec<f64> = Vec::with_capacity(n - lag);
    for i in lag..n - 1 {
        result.push(v[i] - v[i - lag])
    }
    result
}

pub fn autocov(v: &[f64], max_lags: usize) -> Vec<f64> {
    let n = v.len();
    let lags = if max_lags < n {
        max_lags
    }
    else {
        n - 1
    };
    let mut result: Vec<f64> = Vec::with_capacity(lags);
    let mean = stats::mean(v);

    for k in 0..lags {
        let mut sum = 0.0;
        for i in 0..(n - k - 1) {
            sum += (v[i] - mean) * (v[i + k] - mean);
        }
        result.push(sum / n as f64);
    }
    result
}

pub fn acf(v: &[f64], max_lags: usize) -> Vec<f64> {
   let result = autocov(v, max_lags);
   let s0 = result[0];
   result.iter().map(|x| x / s0).collect()
}
