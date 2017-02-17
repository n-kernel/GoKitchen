/**
 * Created by Bram on 17/2/2017.
 */
import React, { Component } from 'react';

import Storage from './Storage';

export default class StorageRow extends Component {

    renderContent() {
        var storages = [];

        for (let storage of this.props.currentStorages) {
            storages.push(
                <Storage id={storage.id} />
            )
        }

        return storages;
    }

    render() {
        return(
            <div className="storage-row">
                <h1>
                    Storages
                </h1>

                { this.renderContent() }
            </div>
        )
    }
}