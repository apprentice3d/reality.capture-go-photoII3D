import React from 'react';
import {H3} from 'hig-react';

class ResultsSection extends React.Component {
    render() {
        const {formats} = this.props;
        console.log(formats);
        return (
            <div>
                <H3>Available results:</H3>
                {formats.map((format) => {
                    return <li>{format}</li>
                })}
            </div>
        )
    }
}

export default ResultsSection;