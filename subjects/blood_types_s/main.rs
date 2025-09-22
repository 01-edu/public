use blood_types_s::*;

fn main() {
    let blood_type = BloodType {
        antigen: Antigen::O,
        rh_factor: RhFactor::Positive,
    };

    println!("recipients of O+ {:?}", blood_type.recipients());
    println!("donors of O+ {:?}", blood_type.donors());

    let another_blood_type = BloodType {
        antigen: Antigen::O,
        rh_factor: RhFactor::Positive,
    };

    println!(
        "donors of O+ can receive from {:?} {}",
        another_blood_type,
        blood_type.can_receive_from(another_blood_type)
    );
}
