# tf_algebra
The purpose of this website is to improve your skills in Linear Algebra 
through the answering of True and False questions. Right now is just a simple website
with nearly 60 True and False questions but the idea is to create an API where this questions
can be fetched. Also has the intend to extend these True and False questions to Calculus or any other subject related to math and computer sciences.

## How it works?
<div>
    <table>
        <tr>
            <th>Resource</th>
            <th>Method</th>
            <th>Description</th>
        </tr>
        <tr>
            <td>/api/questions</td>
            <td>GET</td>
            <td>Queries a list of all available questions.</td>
        </tr>
        <tr>
            <td>/api/users/score</td>
            <td>POST</td>
            <td>Allow to update the score of a given player.</td>
        </tr>
        <tr>
            <td>/api/users</td>
            <td>GET</td>
            <td>Get all availables users in the db</td>
        </tr>
        <tr>
            <td>/api/signUp</td>
            <td>POST</td>
            <td>Returns a time-ordered list of all moves taken during the match.</td>
        </tr>
        <tr>
            <td>/api/login</td>
            <td>POST</td>
            <td>Allow a given user to login</td>
        </tr>
        <tr>
            <td>/api/topic/{name}</td>
            <td>GET</td>
            <td>Retrive a fiven topic by its name</td>
        </tr>
    </table>
</div>

