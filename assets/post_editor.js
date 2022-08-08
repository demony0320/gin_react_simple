class EditorApp extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            category:'',
            subject:'',
            writer:'',
            content:'',
        };
        
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleChange = this.handleChange.bind(this);
    } 

    handleChange(event){
        const target = event.target;
        const name = target.name;
        const value = target.value ;
        this.setState({
            [name]:value
        });
    }

    handleSubmit = async event =>{
        event.preventDefault();
        try { 
            const response = await fetch('/post', {
                method : 'POST',
                headers : {
                 'Content-Type': 'application/json'
                },
                redirect : 'manual',
                body : JSON.stringify(this.state)
            });
            const responseJson = await response.json() 
            console.log(responseJson)
        }catch(error) {
            console.log(error)
        }

        navigation.navigate('/user/list');
    }
    render() {
        return (
                <div className="EditorApp">
                <h2>Using CKEditor 5 build in React</h2>
                    <form onSubmit={this.handleSubmit}>
                    <label>
                        Writer:
                        <input name="writer" type="text" value={this.state.writer} onChange={this.handleChange} />
                    </label>
                    <br />
                    <label>
                        Category:
                        <input name="category" type="text" value={this.state.category} onChange={this.handleChange} />
                    </label>
                    <br />
                    <label>
                        Subject:
                        <input name="subject" type="text" value={this.state.subject} onChange={this.handleChange} />
                    </label>
                    <br />

                        <CKEditor.CKEditor
                        editor={ ClassicEditor }
                        data="<p>Hello from CKEditor 5!</p>"
                        //data={this.state.content}
                        name="content" value={this.state.content} 
                        onReady={ editor => {
                        // You can store the "editor" and use when it is needed.
                        console.log( 'Editor is ready to use!', editor );
                        } }
                        onChange={ ( event, editor ) => {
                        const data = editor.getData();
                        this.setState({['content']:data});
                        console.log( { event, editor, data } );
                        } }
                        onBlur={ ( event, editor ) => {
                        console.log( 'Blur.', editor );
                        } }
                        onFocus={ ( event, editor ) => {
                        console.log( 'Focus.', editor );
                        } }
                        />
                        <input type="submit" value="Submit"/>
                    </form>
                </div>
               );
            }
        }

//export default EditorApp;
const edit= ReactDOM.createRoot(document.getElementById('post_editor'));
edit.render(<EditorApp />);
