/**
 * Created by Bram on 17/2/2017.
 */

const initialState = {
    "currentStorages": [
        {
            "id": 0,
        },
        {
            "id": 1,
        }
    ]
}

export default function reducer(state = initialState, action) {
    switch(action.type) {
        default:
            return state;
    }
}