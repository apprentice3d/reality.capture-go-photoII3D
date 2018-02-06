import React from 'react';
import {H3, Sub1, Button} from 'hig-react';

class ResultsSection extends React.Component {

    constructor(props) {
        super(props);
    }

    getResultsInFormat(format) {
        this.props.getResult(format);
    }


    render() {
        const {formats} = this.props;
        return (
            <div>
                <H3>Scene processing is done</H3>
                <Sub1>Results are avaialble in following formats:</Sub1>
                {formats.map((format) => {
                    return <Button title={format} key={format} type="secondary" onClick={()=> this.getResultsInFormat(format)}/>
                })}
            </div>
        )
    }
}

export default ResultsSection;