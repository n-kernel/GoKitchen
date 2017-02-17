/**
 * Created by Bram on 16/2/2017.
 */
import React, { Component } from 'react';


class Storage extends Component {
    render() {
        return(
            <div key={this.props.id} >
                Storage unit
            </div>
        );
    }
}

export default Storage;