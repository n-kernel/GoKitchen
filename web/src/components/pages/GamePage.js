/**
 * Created by Bram on 16/2/2017.
 */
import React, { Component } from 'react';

import styled from 'styled-components';

import SupplyRow from '../../containers/SupplyRow';
import StorageRow from '../../containers/StorageRow';

export default class GamePage extends Component {

    render() {
        return(
            <div>
                <SupplyRow />
                <StorageRow />
            </div>
        );
    }
}
