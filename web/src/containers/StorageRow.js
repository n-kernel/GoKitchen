/**
 * Created by Bram on 17/2/2017.
 */

import StorageRowComponent from '../components/StorageRow';
import { connect } from 'react-redux';

const mapStateToProps = (state) => {
    return {
        currentStorages: state.storages.currentStorages,
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        
    }
}

const StorageRow = connect(
    mapStateToProps,
    mapDispatchToProps
)(StorageRowComponent)

export default StorageRow;