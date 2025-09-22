use get_document_id::*;

fn main() {
    let office_ok = OfficeOne {
        next_office: Ok(OfficeTwo {
            next_office: Ok(OfficeThree {
                next_office: Ok(OfficeFour {
                    document_id: Ok(13),
                }),
            }),
        }),
    };

    let office_closed = {
        OfficeOne {
            next_office: Ok(OfficeTwo {
                next_office: Err(ErrorOffice::OfficeClose(23)),
            }),
        }
    };

    match office_ok.get_document_id() {
        Ok(id) => println!("Found a document with id {}", id),
        Err(err) => println!("Error: {:?}", err),
    };

    match office_closed.get_document_id() {
        Ok(id) => println!("Found a document with id {}", id),
        Err(err) => println!("Error: {:?}", err),
    };
}
