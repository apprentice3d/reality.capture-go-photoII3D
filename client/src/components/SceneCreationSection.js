import React from 'react'
import {Button, Checkbox, Container, H3, RadioButton, Spacer, Sub2, TextField} from 'hig-react'


class SceneCreationSection extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            scene_name: "some_scene_name",
            output_formats: ["rcm"],
            scene_type: "aerial"
        }
    }


    sceneTypeSetting = (element) => {
        this.setState({
            scene_type: element.target.name
        });
    };

    sceneFormatSetting = (event) => {
        event.stopImmediatePropagation();

        const format = event.target.name;
        const element_index = this.state.output_formats.indexOf(format);
        const {output_formats} = this.state;

        if ( element_index === -1) {
            this.setState({
                output_formats: output_formats.concat(format)
            })
        } else {
            this.setState({
                output_formats: output_formats.slice(0, element_index).concat(output_formats.slice(element_index+1))
            })
        }

    };


    createScene = () => {
      this.props.submit(this.state);
    };

    render() {

        const {scene_name, output_formats, scene_type} = this.state;


        return (
            <div>
                <H3>Scene setup</H3>
                <TextField
                    label="Scene name"
                    placeholder={scene_name}
                    // instructions="set the scene name here"
                />
                <Spacer inset="xxs"/>
                <Container>
                    <Sub2>Output file format:</Sub2>
                    <Checkbox name="rcm"
                              label="rcm: Autodesk ReCap Photo Mesh (default)"
                              onChange={this.sceneFormatSetting}
                              checked={output_formats.indexOf("rcm") > -1}

                    />
                    <Checkbox name="rcs"
                              label="rcs: Autodesk ReCap Point Cloud"
                              onChange={this.sceneFormatSetting}
                              checked={output_formats.indexOf("rcs") > -1}
                    />
                    <Checkbox name="obj"
                              label="obj: Wavefront Object"
                              onChange={this.sceneFormatSetting}
                              checked={output_formats.indexOf("obj") > -1}

                    />
                    <Checkbox name="ortho"
                              label="ortho: Ortho Photo and Elevation Map"
                              onChange={this.sceneFormatSetting}
                              checked={output_formats.indexOf("ortho") > -1}

                    />
                    <Checkbox name="report"
                              label="report: Quality Report"
                              onChange={this.sceneFormatSetting}
                              checked={output_formats.indexOf("report") > -1}

                    />
                </Container>
                <Spacer inset="xxs"/>
                <Container>
                    <Sub2>Scene type:</Sub2>
                    <RadioButton label="aerial: Aerial scene" name="aerial" onChange={this.sceneTypeSetting} checked={scene_type === "aerial"}/>
                    <RadioButton label="object: Object scene" name="object" onChange={this.sceneTypeSetting} checked={scene_type === "object"}/>
                </Container>

                <Spacer inset="xxs"/>
                <div style={{textAlign: "center"}}>
                    <Button title="Create scene" onClick={this.createScene}/>
                </div>
            </div>
        )
    }
}

export default SceneCreationSection;