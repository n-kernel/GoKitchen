/**
 * Created by Bram on 16/2/2017.
 */
import React, { Component } from 'react';

class Supply extends Component {
    render() {
        return(
            <div className="supply-item">
                { this.props.name }
            </div>
        );
    }
}

Supply.propTypes = {
    name: React.PropTypes.string.isRequired,
}

export default Supply;