use crate::ts::random;
use crate::ts::stats;

pub fn sim_gwn(n: usize, ar: &[f64], d: usize, ma: &[f64], std: f64) -> Vec<f64> {
    let ar_len = ar.len();
    let ma_len = ma.len();
    let arma_len = ar_len + ma_len;

    let noise = random::gaussian_white_noise(n + arma_len, std);
    let n_len = noise.len();

    let mut tmp = Vec::with_capacity(n + arma_len);

    //Moving average part
    for i in ma_len..n_len {
        let mut x: f64 = noise[i];
        for j in 0..ma_len {
            x += ma[j] * noise[i - j - 1];
        }
        tmp.push(x)
    }

    for i in arma_len..n + arma_len {
        let mut x = noise[i];
        for j in 0..ar_len {
            x += ar[j] * tmp[i - j - 1];
        }
        tmp.push(x);
    }
    if d > 0 {
        stats::diff(&tmp[ar_len..], d)
    }
    else {
        tmp.drain(ar_len..).collect()
    }
}
