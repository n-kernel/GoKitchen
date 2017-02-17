/**
 * Created by Bram on 17/2/2017.
 */


const initialState = {
    "currentSupplies": [
        {
            "id": 0,
            "name": 'Cheese',
        },
        {
            "id": 1,
            "name": 'Lettuce',
        }
    ]
}

export default function reducer(state = initialState, action) {
    switch(action.type) {
        default:
            return state;
    }
}