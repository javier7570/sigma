use statrs::distribution::Normal;
use rand::distributions::Distribution;
use rand::Rng;


/// Creates a random walk in which x(t)= x(t-1) + w(t), where w(t) 
///  is white noise created using the given random distribution.
///
/// # Arguments
///
/// * `n` - Size of the output series
/// * `d` - Random distribution to create the white noise
/// 
/// # Returns
/// 
/// * A random walk of `n` values using the distribution `d`.
///
pub fn random_walk<D: Distribution<f64>>(n: usize, d: &D) -> Vec<f64> {
    let result = random_vector(n, d);
    let mut x = 0.0;
    result.iter().map(|w| { x += w; x }).collect()
}


/// Creates a random walk in which x(t)= x(t-1) + w(t), where w(t) 
///  is Gaussian whote noise with the given standard deviation.
///
/// # Arguments
///
/// * `n` - Size of the output series
/// * `std` - Standad deviation for the Gaussian distribution
/// 
/// # Returns
/// 
/// * A random walk of `n` values using a Gaussian distribution with standard deviation `std`.
///
pub fn gaussian_random_walk(n: usize, std: f64) -> Vec<f64> {
    let normal = Normal::new(0.0, std).unwrap();
    random_walk(n, &normal)
}


/// Creates a vector of random values using the given distribution.
///
/// # Arguments
///
/// * `n` - Size of the output series
/// * `d` - Distribution to create the random values
/// 
/// # Returns
/// 
/// * A vector of `n` random values using the distribution `d`.
///
pub fn random_vector<D: Distribution<f64>>(n: usize, d: &D) -> Vec<f64> {
    rand::thread_rng().sample_iter(d).take(n).collect()
}


/// Creates a vector of Gaussian white noise values.
///
/// # Arguments
///
/// * `n` - Size of the output series
/// * `std` - Standad deviation for the Gaussian distribution
/// 
/// # Returns
/// 
/// * A vector of `n` Gaussian white noise values with standard deviation `std`.
///
/// # Example
/// 
/// ```
/// use sigma::ts::random;
/// 
/// let v = random::gaussian_white_noise(10000, 1.0);
/// assert_eq!(v.len(), 10000);
/// ```
pub fn gaussian_white_noise(n: usize, std: f64) -> Vec<f64> {
    let normal = Normal::new(0.0, std).unwrap();
    random_vector(n, &normal)
}
