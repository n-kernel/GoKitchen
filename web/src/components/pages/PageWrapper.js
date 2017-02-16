/**
 * Created by Bram on 16/2/2017.
 */
import React, { Component } from 'react';

export default class PageWrapper extends Component {

    render() {
        return(
            <div>

                { this.props.children }

            </div>
        )
    }
}