/**
 * Created by Bram on 16/2/2017.
 */
import React, { Component } from 'react';

import Supply from '../Supply';
import Storage from '../Storage';

export default class GamePage extends Component {


    render() {
        return(
            <div>
                <Supply name="Cheese" />
                <Storage />
            </div>
        );
    }
}
