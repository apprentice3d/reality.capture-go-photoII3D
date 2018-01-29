import React from 'react';
import {Button, Container, H3, IconButton, Tab, Table, Tabs, TextCellContent, TextField} from 'hig-react';





class ImageUploadingSection extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            uri_collection: []
        }
    }


    removeUriFromList = (event, id) => {
        // event.preventDefault();
        const {uri_collection} = this.state;
        let index = -1;
        for (let i = 0; i < uri_collection.length; ++i) {
            if (id === uri_collection[i]["id"]) {
                index = i;
                break;
            }
        }


        if (index >= 0) {
            this.setState({
                uri_collection: uri_collection.slice(0, index).concat(uri_collection.slice(index + 1))
            });
        }
    };


    addNewURI = (value) => {
        const {uri_collection} = this.state;

        const index = uri_collection.length;
        this.setState({
            uri_collection: uri_collection.concat({
                "id": index,
                "image_uri": value
            })
        })
    };


    processURI = (event) => {

        //TODO: add here check if it is valid url
        this.addNewURI(event.target.value);
    };


    uploadImages = () => {
        this.props.upload(this.state.uri_collection);
    };


    render() {

        const column_headers_remote = [
            {
                id: "1",
                HeaderCell: "URI",
                alignment: "left",
                width: "1fr",
                accessor: "image_uri",
                Cell: props => (
                    <TextCellContent text={props.data.image_uri}/>
                )
            },
            {
                id: "2",
                HeaderCell: "",
                alignment: "right",
                width: "0.15fr",
                accessor: "delete",
                Cell: props => (
                    <IconButton title="delete"
                                icon="trash"
                                size="small"
                                onClick={(e) => this.removeUriFromList(e, props.data.id)}/>
                )
            }

        ];




        return (
            <div>
                <H3>Upload images</H3>
                {/*<Container>*/}
                <Tabs
                    // activeTabIndex={this.state.activeTabIndex}
                    // onTabChange={this.setActiveTabIndex}
                >
                    <Tab label="from a Remote Server">
                        <div>
                            <Table
                                // selectable
                                density="compressed"
                                columns={column_headers_remote}
                                data={this.state.uri_collection}
                            />
                        </div>
                    </Tab>
                    <Tab label="from a Local Drive">Activities content</Tab>
                </Tabs>
                <TextField
                    label="add new image uri"
                    // placeholder=""
                    // instructions="add uri and press Enter"
                    onChange={this.processURI}
                />
                {/*</Container>*/}
                <div style={{textAlign: "center"}} >
                    <Button title="Upload images" onClick={this.uploadImages}/>
                </div>
            </div>
        )
    }
}

export default ImageUploadingSection;