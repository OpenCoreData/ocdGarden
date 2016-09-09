import graphene

schema = graphene.Schema()

# This will be our root query
class Query(graphene.ObjectType):
    username = graphene.StringField(description='The username')

    def resolve_username(self, *args):
        return 'Hello World'

# Here we set the root query for our schema
schema.query = Query

result = schema.execute('{ username }')

# result.data should be {'username': 'Hello World'}
username = result.data['username']

print(username)

