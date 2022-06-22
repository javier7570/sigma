pub fn mean(v: &[f64]) -> f64
{
    v.iter().sum::<f64>() / v.len() as f64
}

pub fn variance(v: &[f64]) -> f64
{
    let mean = mean(v);
    v.iter().fold(0.0, |sum, val| sum + (val - mean).powi(2)) / v.len() as f64
}

fn moving(v: &[f64], size: usize, fun: impl Fn(&[f64]) -> f64) -> Vec<f64> {
    let mut result: Vec<f64>;
    let len = v.len();
    if len >= size && size != 0 {
         result = Vec::with_capacity(len - size + 1);

         let mut index = 0;
         for win in v.windows(size) {
            result[index] = fun(win);
            index += 1;
         }
    }
    else {
        result = Vec::with_capacity(0);
    }
    result
}

pub fn moving_mean(v: &[f64], size: usize) -> Vec<f64> {
    moving(v, size, mean)
}

pub fn moving_var(v: &[f64], size: usize) -> Vec<f64> {
    moving(v, size, variance)
}

pub fn diff(v: &[f64], lag: usize) -> Vec<f64> {
    let n = v.len();
    if n >= lag {
        let mut result = Vec::with_capacity(n - lag);
        for i in 0..n - lag {
            result.push(v[i + lag] - v[i]);
        }
        result
    }
    else {
        Vec::with_capacity(0)
    }
}

pub fn diff_emplace(v: &mut Vec<f64>, lag: usize)  {
    let n = v.len();
    if n >= lag {
        for i in 0..n - lag {
            v[i] = v[i + lag] - v[i];
        }
        v.truncate(n - lag);
    }
    else {
        v.clear();
    }
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
    let mean = mean(v);

    for k in 0..lags {
        let mut sum = 0.0;
        for i in 0..(n - k - 1) {
            sum += (v[i] - mean) * (v[i + k] - mean);
        }
        result.push(sum / n as f64);
    }
    result
}


/// Calculates the autocorrelation function for the given time series
///
/// # Arguments
///
/// * `v` - Time series
/// * `max_lags` - Standad deviation for the Gaussian distribution
/// 
/// #Returns
/// 
/// * A vector of n random values using the distribution d.
///
pub fn acf(v: &[f64], max_lags: usize) -> Vec<f64> {
   let result = autocov(v, max_lags);
   let s0 = result[0];
   result.iter().map(|x| x / s0).collect()
}
