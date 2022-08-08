class UserForm extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            userId:'',
            name:'',
            password:'',
        };
        
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    } 
    
    handleChange(event){
        const target = event.target;
        const value = target.value;
        const name = target.name;
        this.setState({
            [name]:value
        });
    }

    handleSubmit = async event =>{
        event.preventDefault();
        try { 
            const response = await fetch('/user', {
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

    render(){
        return (
            <form onSubmit={this.handleSubmit}>
                <label>
                    Id:
                    <input name="userId" type="text"  data={this.state.userId} onChange={this.handleChange} />
                </label>
                <br />
                <label>
                    Name:
                    <input name="name" type="text"  data={this.state.name} onChange={this.handleChange} />
                </label>
                <br />
                <label>
                    Password:
                    <input name="password" type="text"  data={this.state.password} onChange={this.handleChange} />
                </label>
                <input type="submit" value="Submit"/>
            </form>
        );
    }
}

const user= ReactDOM.createRoot(document.getElementById('userform'));
user.render(<UserForm />);

