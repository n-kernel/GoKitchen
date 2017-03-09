/**
 * Created by Bram on 17/2/2017.
 */

import SupplyRowComponent from '../components/SupplyRow';
import { connect } from 'react-redux';

const mapStateToProps = (state) => {
    return {
        currentSupplies: state.supplies.currentSupplies,
    }
}

const mapDispatchToProps = (dispatch) => {
    return {

    }
}

const SupplyRow = connect(
    mapStateToProps,
    mapDispatchToProps
)(SupplyRowComponent)

export default SupplyRow;