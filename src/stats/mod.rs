pub fn mean(v: &[f64]) -> f64
{
    v.iter().sum::<f64>() / v.len() as f64
}

pub fn variance(v: &[f64]) -> f64
{
    let mean = mean(v);
    v.iter().fold(0.0, |sum, val| sum + (val - mean).powi(2)) / v.len() as f64
}
