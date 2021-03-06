import React, {Component} from 'react'
import {withRouter, WithRouterProps} from 'react-router'
import {connect} from 'react-redux'

// Components
import {ErrorHandling} from 'src/shared/decorators/errors'
import OrgTabbedPage from 'src/organizations/components/OrgTabbedPage'
import OrgHeader from 'src/organizations/components/OrgHeader'
import SettingsTabbedPage from 'src/settings/components/SettingsTabbedPage'
import SettingsHeader from 'src/settings/components/SettingsHeader'
import {Page} from '@influxdata/clockface'
import GetResources from 'src/resources/components/GetResources'
import Members from 'src/members/components/Members'
import {FeatureFlag} from 'src/shared/utils/featureFlag'

// Utils
import {pageTitleSuffixer} from 'src/shared/utils/pageTitles'
import {getByID} from 'src/resources/selectors'

// Types
import {AppState, Organization, ResourceType} from 'src/types'

interface StateProps {
  org: Organization
}

type Props = WithRouterProps & StateProps

@ErrorHandling
class MembersIndex extends Component<Props> {
  constructor(props) {
    super(props)
  }

  public render() {
    const {org, children} = this.props

    return (
      <>
        <FeatureFlag name="treeNav">
          <Page titleTag={pageTitleSuffixer(['Members', 'Organization'])}>
            <OrgHeader />
            <OrgTabbedPage activeTab="members" orgID={org.id}>
              <GetResources resources={[ResourceType.Members]}>
                <Members />
              </GetResources>
            </OrgTabbedPage>
          </Page>
        </FeatureFlag>
        <FeatureFlag name="treeNav" equals={false}>
          <Page titleTag={pageTitleSuffixer(['Members', 'Settings'])}>
            <SettingsHeader />
            <SettingsTabbedPage activeTab="members" orgID={org.id}>
              <GetResources resources={[ResourceType.Members]}>
                <Members />
              </GetResources>
            </SettingsTabbedPage>
          </Page>
        </FeatureFlag>
        {children}
      </>
    )
  }
}

const mstp = (state: AppState, props: Props) => {
  const org = getByID<Organization>(
    state,
    ResourceType.Orgs,
    props.params.orgID
  )

  return {
    org,
  }
}

export default connect<StateProps>(
  mstp,
  null
)(withRouter<{}>(MembersIndex))
