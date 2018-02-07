import React from 'react';
import {Button, Container, H3, IconButton, Spacer, Tab, Table, Tabs, TextCellContent, TextField} from 'hig-react';

import DropzoneComponent from 'react-dropzone-component'


const sample_images = [
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5120.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5121.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5122.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5123.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5124.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5125.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5126.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5127.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5128.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5129.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5130.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5131.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5132.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5133.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5134.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5135.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5136.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5137.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5138.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5139.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5140.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5141.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5142.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5143.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5144.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5145.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5146.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5147.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5148.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5149.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5150.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5151.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5152.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5153.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5154.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5155.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5156.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5157.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5158.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5159.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5160.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5161.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5162.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5163.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5164.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5165.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5166.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5167.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5168.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5169.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5170.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5171.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5172.JPG",
    // "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5173.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5174.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5175.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5176.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5177.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5178.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5179.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5180.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5181.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5182.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5183.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5184.JPG",
    // "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5185.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5186.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5187.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5188.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5189.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5190.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5191.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5192.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5193.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5194.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5195.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5196.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5197.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5198.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5199.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5200.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5201.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5202.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5203.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5204.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5205.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5206.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5207.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5208.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5209.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5210.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5211.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5212.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5213.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5214.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5215.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5216.JPG"
];


let dropZoneConfig = {
    iconFiletypes: ['.jpg'],
    showFiletypeIcon: true,
    postUrl: "http://" + window.location.hostname + ":3000/uploadLocalImages"
};


class ImageUploadingSection extends React.Component {
    constructor(props) {
        super(props);

        let sample = [];
        // populate example with sample images
        for (let i = 0; i < sample_images.length; ++i) {
            sample.push({
                "id": i,
                "image_uri": sample_images[i]
            })
        }

        this.state = {
            uri_collection: sample,
            done_uploading: false
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
        this.setState({
            done_uploading: true
        });

        this.props.upload(this.state.uri_collection);
    };


    doneAddingLocalImages = () => {
        this.setState({
            done_uploading: true
        });
        this.props.upload([]);
    };


    render() {

        const {done_uploading} = this.state;

        let visibility = done_uploading ? "none" : "inline";

        const column_headers_remote = [
            {
                id: "1",
                HeaderCell: "URI",
                alignment: "left",
                width: "1fr",
                accessor: "image_uri",
                Cell: props => (
                    <TextCellContent text={props.data.image_uri} detail={props.data.image_uri}/>

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
                        <div style={{display: visibility}}>
                            <Table
                                // selectable
                                density="compressed"
                                columns={column_headers_remote}
                                data={this.state.uri_collection}

                            />

                            <TextField
                                label="add new image uri"
                                // placeholder=""
                                // instructions="add uri and press Enter"
                                onChange={this.processURI}
                            />
                        </div>
                        <Spacer inset="xxs"/>
                        <div style={{textAlign: "center"}}>
                            <Button title="Upload images" onClick={this.uploadImages} disabled={done_uploading}/>
                        </div>
                    </Tab>
                    <Tab label="from a Local Drive">
                        <Container>
                            <div style={{display: visibility}}>
                                <DropzoneComponent
                                    config={dropZoneConfig}
                                    // eventHandlers={(file) => this.prepareImageForUploading(file)}
                                    // djsConfig={{autoProcessQueue:false}}
                                />
                            </div>
                            <div style={{textAlign: "center"}}>
                                <Spacer inset="xxs"/>
                                <Button title="Done" onClick={this.doneAddingLocalImages} disabled={done_uploading}/>

                            </div>
                        </Container>
                    </Tab>
                </Tabs>


            </div>
        )
    }
}

export default ImageUploadingSection;