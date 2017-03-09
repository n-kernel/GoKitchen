/**
 * Created by Bram on 17/2/2017.
 */
import React, { Component } from 'react';
import Supply from './Supply';

export default class SupplyRow extends Component {

    renderContent() {
        var supplies = [];

        for (let supply of this.props.currentSupplies) {
            supplies.push(
                <Supply id={supply.id} name={supply.name} />
            )
        }
        
        return supplies;
    }

    render() {
        return(
            <div className="supply-row">
                <h1>
                    Supplies
                </h1>
                { this.renderContent() }
            </div>
        );
    }
}