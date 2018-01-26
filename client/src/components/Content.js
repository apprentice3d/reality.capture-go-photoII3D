import React from 'react'
import {Container, Button, H3} from 'hig-react'

import { VerticalTimeline, VerticalTimelineElement }  from 'react-vertical-timeline-component';
import 'react-vertical-timeline-component/style.min.css';


class Content extends React.Component {
    render() {
        return (

            <Container>
                <Container >
                    Setup
                </Container>

            <VerticalTimeline>
                <VerticalTimelineElement
                    className="scene_creation_element"
                    date="Scene creation"
                    iconStyle={{ background: '#0696D7', color: '#fff' }}
                    // icon={<WorkIcon />}
                >
                    <H3>Create a scene</H3>
                    <Button title="create scene "/>
                </VerticalTimelineElement>

                <VerticalTimelineElement
                    className="vertical-timeline-element--work"
                    date="Resource uploading"
                    iconStyle={{ background: '#0696D7', color: '#666666' }}
                    // icon={<WorkIcon />}
                >
                    <H3>Upload images</H3>
                </VerticalTimelineElement>

                <VerticalTimelineElement
                    className="vertical-timeline-element--work"
                    date="Initiate work"
                    iconStyle={{ background: '#0696D7', color: '#666666' }}
                    // icon={<WorkIcon />}
                >
                    <H3>Start scene processing</H3>
                </VerticalTimelineElement>

                <VerticalTimelineElement
                    className="vertical-timeline-element--work"
                    date="Processing"
                    iconStyle={{ background: '#0696D7', color: '#666666' }}
                    // icon={<WorkIcon />}
                >
                    <H3>Query work status</H3>
                </VerticalTimelineElement>

                <VerticalTimelineElement
                    className="vertical-timeline-element--work"
                    date="Result query"
                    iconStyle={{ background: '#0696D7', color: '#666666' }}
                    // icon={<WorkIcon />}
                >
                    <H3>Get results</H3>
                </VerticalTimelineElement>
                {/*<VerticalTimelineElement*/}
                    {/*className="vertical-timeline-element--work"*/}
                    {/*date="2008 - 2010"*/}
                    {/*iconStyle={{ background: 'rgb(33, 150, 243)', color: '#fff' }}*/}
                    {/*// icon={<WorkIcon />}*/}
                {/*>*/}
                    {/*<h3 className="vertical-timeline-element-title">Web Designer</h3>*/}
                    {/*<h4 className="vertical-timeline-element-subtitle">Los Angeles, CA</h4>*/}
                    {/*<p>*/}
                        {/*User Experience, Visual Design*/}
                    {/*</p>*/}
                {/*</VerticalTimelineElement>*/}
                {/*<VerticalTimelineElement*/}
                    {/*className="vertical-timeline-element--work"*/}
                    {/*date="2006 - 2008"*/}
                    {/*iconStyle={{ background: 'rgb(33, 150, 243)', color: '#fff' }}*/}
                    {/*// icon={<WorkIcon />}*/}
                {/*>*/}
                    {/*<h3 className="vertical-timeline-element-title">Web Designer</h3>*/}
                    {/*<h4 className="vertical-timeline-element-subtitle">San Francisco, CA</h4>*/}
                    {/*<p>*/}
                        {/*User Experience, Visual Design*/}
                    {/*</p>*/}
                {/*</VerticalTimelineElement>*/}
                {/*<VerticalTimelineElement*/}
                    {/*className="vertical-timeline-element--education"*/}
                    {/*date="April 2013"*/}
                    {/*iconStyle={{ background: 'rgb(233, 30, 99)', color: '#fff' }}*/}
                    {/*// icon={<SchoolIcon />}*/}
                {/*>*/}
                    {/*<h3 className="vertical-timeline-element-title">Content Marketing for Web, Mobile and Social Media</h3>*/}
                    {/*<h4 className="vertical-timeline-element-subtitle">Online Course</h4>*/}
                    {/*<p>*/}
                        {/*Strategy, Social Media*/}
                    {/*</p>*/}
                {/*</VerticalTimelineElement>*/}
                {/*<VerticalTimelineElement*/}
                    {/*className="vertical-timeline-element--education"*/}
                    {/*date="November 2012"*/}
                    {/*iconStyle={{ background: 'rgb(233, 30, 99)', color: '#fff' }}*/}
                    {/*// icon={<SchoolIcon />}*/}
                {/*>*/}
                    {/*<h3 className="vertical-timeline-element-title">Agile Development Scrum Master</h3>*/}
                    {/*<h4 className="vertical-timeline-element-subtitle">Certification</h4>*/}
                    {/*<p>*/}
                        {/*Creative Direction, User Experience, Visual Design*/}
                    {/*</p>*/}
                {/*</VerticalTimelineElement>*/}
                {/*<VerticalTimelineElement*/}
                    {/*className="vertical-timeline-element--education"*/}
                    {/*date="2002 - 2006"*/}
                    {/*iconStyle={{ background: 'rgb(233, 30, 99)', color: '#fff' }}*/}
                    {/*// icon={<SchoolIcon />}*/}
                {/*>*/}
                    {/*<h3 className="vertical-timeline-element-title">Bachelor of Science in Interactive Digital Media Visual Imaging</h3>*/}
                    {/*<h4 className="vertical-timeline-element-subtitle">Bachelor Degree</h4>*/}
                    {/*<p>*/}
                        {/*Creative Direction, Visual Design*/}
                    {/*</p>*/}
                {/*</VerticalTimelineElement>*/}
            </VerticalTimeline>
<Container>
    <H3>Finish</H3>
</Container>
            </Container>
        )
    }
}

export default Content;