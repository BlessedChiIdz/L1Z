fn main() {
    let a:i32 = 2;
    let p:i32 = 25;
    let y:i32 = 9;
    let ret:i32 = big_guy(a, p, y);
    println!("unswer is: {}", ret);
}

fn big_guy(a:i32, p:i32, y:i32) -> i32{
    let m:i32 = 6;
    let k = p/m + 1;


    
    
    let mut v: Vec<i32> = Vec::new();
    let mut v1: Vec<i32> = Vec::new();
    for i in 0..(m-1){
        let temp:i32 = (i32::pow(a, i as u32) * y)%p;
        v.push(temp);
    }
    for i in 0..k+1{
        let temp:i32 = (i32::pow(a,(m*i) as u32))%p;
        v1.push(temp);
    }
    let mut res = (-1,-1);
    'outer: for i in 0..(m-1){
        let temp:i32 = v[i as usize];
        for j in 0..k{
            if temp == v1[j as usize] {
                res = (i , j);
                break 'outer;
            }
        }
    }
    println!("{:?}", v);
    println!("{:?}", v1);
    println!("{:?}", res);
    if res == (-1, -1){
        panic!("Ne validno")
    }
    let unsw: i32 = res.0 * m - res.1;
    return unsw;
}