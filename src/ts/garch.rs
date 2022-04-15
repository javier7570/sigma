use crate::ts::random;

pub fn sim_gwn(n: usize, p: &[f64], q: &[f64]) -> Vec<f64> {
    let pq_max = if p.len() > q.len() { p.len() }
                  else { q.len() };
    let n_total = n + pq_max;

    let noise = random::gaussian_white_noise(n_total, 1.0);
              
    let mut eps = vec![0.0; n_total];
    let mut sigsq = vec![0.0; n_total];
    
    for i in pq_max..n_total {
        let mut x = p[0];
        for j in 1..p.len() {
            x += p[j] * eps[i - j - 1] * eps[i - j - 1];
        }
        for j in 1..p.len() {
            x += q[j] * sigsq[i - j - 1];
        }
        sigsq[i] = x;
        eps[i] = noise[i] * x.sqrt();
    }

    eps.drain(0..pq_max).collect()
}
